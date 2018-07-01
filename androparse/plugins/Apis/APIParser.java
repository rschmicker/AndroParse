import java.util.ArrayList;
import com.unh.unhcfreg.RapidAndroidParser;
import com.unh.unhcfreg.QueryBlock;
import dataComponent.MethodElement;

public class APIParser {
	public static void main(String args[]){

		final RapidAndroidParser rapid = new RapidAndroidParser();
		String apk = args[0];
		rapid.setSingleApk(apk);
		 rapid.setQuery(new QueryBlock(){

			public void queries() {

				// print API list
				ArrayList<MethodElement>apiList=rapid.getApiList();
				for(int j = 0; j<apiList.size(); j++){
					//apiList.get(j).printFields();
					if(apiList.get(j).returnValueType==null){
						apiList.get(j).returnValueType="";
					}
					if(apiList.get(j).returnValueType==null){
						apiList.get(j).className="";
					}
					if(apiList.get(j).returnValueType==null){
						apiList.get(j).methodName="";
					}
					System.out.print(apiList.get(j).returnValueType+" "+apiList.get(j).className+"."+apiList.get(j).methodName+" ");
					System.out.print("(");
					if(apiList.get(j).hasParameter()){

						for(int k=0;k<apiList.get(j).parameterType.length;k++){
							System.out.print(apiList.get(j).parameterType[k]);
							if(j!=(apiList.get(j).parameterType.length-1)){
								System.out.print(", ");
							}
						}

					}
					System.out.println(")");
				}
			}
		});
	}
}