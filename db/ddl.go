package db

const CreateUsersTable = `Create table if not exists users(
	id integer primary key autoincrement,
	name text not null,
	surname text not null,
	age integer not null,
	gender text not null,
	login text not null,
	password text not null,
	role text not null,
	remove boolean not null default false
);`

const CreateAccountsTable = `Create table if not exists accounts(
	id integer primary key autoincrement,
	user_id int64 references users(id),
	amount int64,
	number text,
	currency text,
	remove boolean default false
);`

const CreateATMsTable = `Create table if not exists atms(
	id integer primary key autoincrement,
	address text not null,
	status boolean default true
);`

const TransactionsHistory = `Create table if not exists (
	id integer primary key autoincrement,
	user_id integer references users(id),
	type 
);`