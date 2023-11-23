
#[tokio::main]
async fn main() {

    let http_proxy = reqwest::Proxy::http("http://gw-us.nstproxy.com:24125")
        .unwrap().basic_auth("user", "password");

    let client = reqwest::ClientBuilder::new().proxy(http_proxy).build().unwrap();

    let resp = client.get("API_URL").send().await.unwrap();

    println!("{:?}", resp.text().await.unwrap());

}
