import React from 'react';
import './App.css';
import { withAuthenticator} from '@aws-amplify/ui-react';
import { Amplify, Auth } from 'aws-amplify';
import awsconfig from './aws-exports';


import '@aws-amplify/ui-react/styles.css';

Amplify.configure(awsconfig);


function App() {
   return (
    <div>
      <h1>Help!</h1>
      
    </div>
   );
}

export default withAuthenticator(App);