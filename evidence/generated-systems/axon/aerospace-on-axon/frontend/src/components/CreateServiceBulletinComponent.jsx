import React, { Component } from 'react'
import ServiceBulletinService from '../services/ServiceBulletinService';

class CreateServiceBulletinComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                bulletinNumber: '',
                category: ''
        }
        this.changebulletinNumberHandler = this.changebulletinNumberHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ServiceBulletinService.getServiceBulletinById(this.state.id).then( (res) =>{
                let serviceBulletin = res.data;
                this.setState({
                    bulletinNumber: serviceBulletin.bulletinNumber,
                    category: serviceBulletin.category
                });
            });
        }        
    }
    saveOrUpdateServiceBulletin = (e) => {
        e.preventDefault();
        let serviceBulletin = {
                serviceBulletinId: this.state.id,
                bulletinNumber: this.state.bulletinNumber,
                category: this.state.category
            };
        console.log('serviceBulletin => ' + JSON.stringify(serviceBulletin));

        // step 5
        if(this.state.id === '_add'){
            serviceBulletin.serviceBulletinId=''
            ServiceBulletinService.createServiceBulletin(serviceBulletin).then(res =>{
                this.props.history.push('/serviceBulletins');
            });
        }else{
            ServiceBulletinService.updateServiceBulletin(serviceBulletin).then( res => {
                this.props.history.push('/serviceBulletins');
            });
        }
    }
    
    changebulletinNumberHandler= (event) => {
        this.setState({bulletinNumber: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/serviceBulletins');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ServiceBulletin</h3>
        }else{
            return <h3 className="text-center">Update ServiceBulletin</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> bulletinNumber:&emsp; </label>
                                                <input placeholder="bulletinNumber" name="bulletinNumber" className="form-control" value={this.state.bulletinNumber} onChange={this.changebulletinNumberHandler}/>

                                            <label> Category:&emsp; </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          Recommended
                      </option>
                      <option name="Category" className="form-control" >
                          Optional
                      </option>
                      <option name="Category" className="form-control" >
                          Alert
                      </option>
                      <option name="Category" className="form-control" >
                          Mandatory
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateServiceBulletin}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateServiceBulletinComponent
