import ner as n
import sqliteconnector as s


if __name__ == "__main__":
    db_name = "mediums"
    articles = s.get_texts(db_name)
    articles_lemmas = n.articles_ner_and_count(articles)
    if len(articles_lemmas):
        s.persist_articles_lemmas(articles_lemmas, db_name)
