import sqlite3


def get_texts(db_name, ids):
    con = sqlite3.connect(f'../../{db_name}.db')
    texts = con.execute(f"SELECT id, text FROM articles WHERE id IN {ids}").fetchall()
    con.close()
    return texts


def persist_articles_lemmas(articles_lemmas, db_name):
    con = sqlite3.connect(f'../../{db_name}.db')
    for article_id in articles_lemmas.keys():
        lemma, count = articles_lemmas[article_id]
        con.execute(f"INSERT INTO articles_lemmas (article_id, lemma, count)"
                    f"VALUES ({article_id}, {lemma}, {count})")
    con.close()
