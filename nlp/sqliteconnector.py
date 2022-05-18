import sqlite3


def get_texts(db_name):
    print("retrieving text from the SQLite database")
    con = sqlite3.connect(f'../database/{db_name}.db')
    texts = con.execute(f"SELECT id, text FROM articles WHERE id NOT IN (SELECT article_id FROM "
                        f"articles_lemmas)").fetchall()
    con.close()
    return texts


def persist_articles_lemmas(articles_lemmas, db_name):
    con = sqlite3.connect(f'../database/{db_name}.db')
    for article_id in articles_lemmas.keys():
        for lemma_type in articles_lemmas[article_id]:
            count = articles_lemmas[article_id][lemma_type]
            try:
                lemma, type = lemma_type[:-4], lemma_type[-3:]
                con.execute(f"INSERT INTO articles_lemmas (article_id, lemma, type, count) VALUES "
                            f"({article_id}, '{lemma}', '{type}' , {count})")
            except:
                print(f"error while trying insert lemma, type and count {lemma_type} for article ID"
                      f" {article_id}")
                continue
    con.commit()
    con.close()
