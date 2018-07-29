import { Injectable } from '@angular/core';
import { CognitoUserPool, CognitoUser, AuthenticationDetails} from 'amazon-cognito-identity-js';
import { environment } from '../../environments/environment';

@Injectable()
export class CognitoService {

  public static CONFIG = {
    UserPoolId: environment.userPoolId,
    ClientId: environment.clientId
  }

  auth(username, password, callback){
    let user = new CognitoUser({
      Username: username,
      Pool: this.getUserPool()
    })

    let authDetails = new AuthenticationDetails({
      Username: username,
      Password: password
    })

    user.authenticateUser(authDetails, {
      onSuccess: res => {
        callback(null, res.getIdToken().getJwtToken())
      },
      onFailure: err => {
        callback(err, null)
      }
    })
  }

  getUserPool() {
    return new CognitoUserPool(CognitoService.CONFIG);
  }

  getCurrentUser() {
    return this.getUserPool().getCurrentUser();
  }

}
