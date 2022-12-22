
use tokoku;

-- DROP
DROP TABLE transaksi_barang ;
DROP TABLE transaksi;
DROP TABLE barang;
DROP TABLE customer;
DROP TABLE pegawai;


-- CREATE
create table pegawai (
    id_pegawai int not null AUTO_INCREMENT,
    username varchar(100),
    password varchar(100),
    primary key(id_pegawai)
);



create table customer (
    id_customer int not null AUTO_INCREMENT,
    id_pegawai int not null,
    username varchar(100),
    primary key(id_customer),
    foreign key(id_pegawai) references pegawai(id_pegawai)
);

ALTER TABLE tokoku.customer DROP FOREIGN KEY customer_ibfk_1;
ALTER TABLE tokoku.customer ADD CONSTRAINT customer_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) ON DELETE CASCADE ON UPDATE CASCADE;


create table transaksi (
    id_transaksi int not null AUTO_INCREMENT,
    id_pegawai int not null,
    id_customer int not null,
    created_date timestamp not null default current_timestamp,
    primary key(id_transaksi),
    foreign key(id_pegawai) references pegawai(id_pegawai),
    foreign key(id_customer) references customer(id_customer)
);

ALTER TABLE tokoku.transaksi DROP FOREIGN KEY transaksi_ibfk_1;
ALTER TABLE tokoku.transaksi ADD CONSTRAINT transaksi_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE tokoku.transaksi DROP FOREIGN KEY transaksi_ibfk_2;
ALTER TABLE tokoku.transaksi ADD CONSTRAINT transaksi_ibfk_2 FOREIGN KEY (id_customer) REFERENCES tokoku.customer(id_customer) ON DELETE CASCADE ON UPDATE CASCADE;


create table barang (
    id_barang int not null AUTO_INCREMENT,
    id_pegawai int not null,
    nama varchar(100),
    quantity int,
    created_date timestamp not null default current_timestamp,
    primary key(id_barang),
    foreign key(id_pegawai) references pegawai(id_pegawai)
);

ALTER TABLE tokoku.barang DROP FOREIGN KEY barang_ibfk_1;
ALTER TABLE tokoku.barang ADD CONSTRAINT barang_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) ON DELETE CASCADE ON UPDATE CASCADE;


create table transaksi_barang (
    id_transaksi int not null,
    id_barang int not null,
    foreign key(id_transaksi) references transaksi(id_transaksi),
    foreign key(id_barang) references barang(id_barang),
    PRIMARY KEY(id_transaksi, id_barang)
);

-- ALTER 
ALTER TABLE tokoku.customer DROP FOREIGN KEY customer_ibfk_1;
ALTER TABLE tokoku.customer 
ADD CONSTRAINT customer_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) 
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE tokoku.barang DROP FOREIGN KEY barang_ibfk_1;
ALTER TABLE tokoku.barang 
ADD CONSTRAINT barang_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) 
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE tokoku.transaksi DROP FOREIGN KEY transaksi_ibfk_1;
ALTER TABLE tokoku.transaksi 
ADD CONSTRAINT transaksi_ibfk_1 FOREIGN KEY (id_pegawai) REFERENCES tokoku.pegawai(id_pegawai) 
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE tokoku.transaksi DROP FOREIGN KEY transaksi_ibfk_2;
ALTER TABLE tokoku.transaksi 
ADD CONSTRAINT transaksi_ibfk_2 FOREIGN KEY (id_customer) REFERENCES tokoku.customer(id_customer) 
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE tokoku.transaksi_barang DROP FOREIGN KEY transaksi_barang_ibfk_1;
ALTER TABLE tokoku.transaksi_barang 
ADD CONSTRAINT transaksi_barang_ibfk_1 FOREIGN KEY (id_transaksi) REFERENCES tokoku.transaksi(id_transaksi) 
ON DELETE CASCADE ON UPDATE CASCADE;


ALTER TABLE tokoku.transaksi_barang DROP FOREIGN KEY transaksi_barang_ibfk_2;
ALTER TABLE tokoku.transaksi_barang 
ADD CONSTRAINT transaksi_barang_ibfk_2 FOREIGN KEY (id_barang) REFERENCES tokoku.barang(id_barang) 
ON DELETE CASCADE ON UPDATE CASCADE;


