import sqlite3

def get_texts(db_name):
    con = sqlite3.connect(f'../../{db_name}.db')
    texts = con.execute("select text from articles").fetchall()
    con.close()
    return [t[0] for t in texts]