import { Component } from '@angular/core';
import { CognitoUserPool, CognitoUser} from 'amazon-cognito-identity-js';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  public _POOL_DATA: any = {
    UserPoolId: "us-east-1_qTCypPvBf",
    ClientId: "6ndbtp2jso02odhjo2i1i103n0"
};

  constructor(){
  }

  

  getUserPool() {
    return new CognitoUserPool(this._POOL_DATA);
  }

  getCurrentUser() {
    return this.getUserPool().getCurrentUser();
  }
}
