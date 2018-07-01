import java.util.ArrayList;
import com.unh.unhcfreg.RapidAndroidParser;
import com.unh.unhcfreg.QueryBlock;
import dataComponent.StringElement;



public class StringParser {
	public static void main(final String args[]){
		
		final RapidAndroidParser rapid = new RapidAndroidParser();
		String apk = args[0];
		 rapid.setSingleApk(apk);
		 rapid.setQuery(new QueryBlock(){
		 
			public void queries() {
				//print sting list
				ArrayList<StringElement>stringList=rapid.getStringList();
				for(int i = 0; i< stringList.size(); i ++){
					System.out.println(""+stringList.get(i).stringContent);
				}				
			}
		 });
	}
}