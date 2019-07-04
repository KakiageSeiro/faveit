package model

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

//twiiterAPI用パラメータ
type param struct {
	user_id 			string	//対象ユーザーのユーザーID。
	screen_name 		string	//対象ユーザーのスクリーンネーム。
	count 				uint8	//結果の数。1〜200の間で指定する。
	since_id 			uint64	//ページングに利用する。ツイートのIDを指定すると、これを含まず、これより未来のツイートを取得できる。
	max_id 				uint64	//ページングに利用する。ツイートのIDを指定すると、これを含まず、これより過去のツイートを取得できる。
	include_entities 	bool	//ツイートオブジェクト内のentitiesプロパティを含めるか否か。
}

//twetterAPIから取得する結果
type fav struct{


}

const requestUrl = "https://api.twitter.com/1.1/favorites/list.json"	// エンドポイント"

//いいね一覧を取得
func requestFavList(bearerToken string)([]fav, error){

	//仮のパラメータ
	param := param{
		user_id:"onikakusi",
		screen_name:"",
		count:10,
		since_id:nil,
		max_id:nil,
		include_entities:true,
	}

	//URLを取得
	u, err := retrieveUrl(param)
	if err != nil {
		return nil, err
	}

	//リクエスト実行
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	//結果のjsonをパース



	

	return resp, nil
}

//リクエスト用URLを取得
func retrieveUrl(param param)(*url.URL, error){
	//URLの構造体にパース
	u, err := url.Parse(requestUrl)
	if err != nil {
		return nil, err
	}
	//パラメータをURLに連結してリクエスト実行
	q := u.Query()
	setParam(param, q)

	u.RawQuery = q.Encode()

	return u, nil
}


//URLにパラメータをセット
func setParam(param param, q url.Values) {

	//reflect.Indirectでポインタが指し示す値を取得
	//reflect.ValueOfで構造体インスタンスの値を取得
	value := reflect.Indirect(reflect.ValueOf(param))

	//型を取得
	t := value.Type()

	//型のフィールド数ループ
	for i := 0; i < t.NumField(); i++ {
		//フィールド名
		println("Field: " + t.Field(i).Name)

		//値
		field := value.Field(i)
		fieldInterface := field.Interface()
		if intValue, ok := fieldInterface.(int); ok {
			//数値型の場合
			println("Value: " + strconv.Itoa(intValue))
			//パラメータを追加
			q.Set(t.Field(i).Name, strconv.Itoa(intValue))
		} else {
			//文字列型の場合
			println("Value: " + field.String())
			//パラメータを追加
			q.Set(t.Field(i).Name, field.String())
		}
	}
}

