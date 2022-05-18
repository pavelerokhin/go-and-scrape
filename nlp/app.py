import ner as n
import sqliteconnector as s

from flask import Flask
from flask_cors import CORS

app = Flask(__name__)
CORS(app)

@app.route("/")
def root():
    return {"message": "ok"}, 200

@app.route("/ner")
def ner():
    ner_manager()
    return {"message": "ok"}, 200




def ner_manager():
    print("starting NER")
    db_name = "mediums"
    articles = s.get_texts(db_name)
    articles_lemmas = n.articles_ner_and_count(articles)
    if len(articles_lemmas):
        print("entities have been recognized in", len(articles_lemmas), "articles")
        s.persist_articles_lemmas(articles_lemmas, db_name)
    else:
        print("no articles contain names entities")

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=7070)
