import React, { Component } from 'react'
import ServiceBulletinService from '../services/ServiceBulletinService';

class UpdateServiceBulletinComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                bulletinNumber: '',
                category: ''
        }
        this.updateServiceBulletin = this.updateServiceBulletin.bind(this);

        this.changebulletinNumberHandler = this.changebulletinNumberHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    componentDidMount(){
        ServiceBulletinService.getServiceBulletinById(this.state.id).then( (res) =>{
            let serviceBulletin = res.data;
            this.setState({
                bulletinNumber: serviceBulletin.bulletinNumber,
                category: serviceBulletin.category
            });
        });
    }

    updateServiceBulletin = (e) => {
        e.preventDefault();
        let serviceBulletin = {
            serviceBulletinId: this.state.id,
            bulletinNumber: this.state.bulletinNumber,
            category: this.state.category
        };
        console.log('serviceBulletin => ' + JSON.stringify(serviceBulletin));
        console.log('id => ' + JSON.stringify(this.state.id));
        ServiceBulletinService.updateServiceBulletin(serviceBulletin).then( res => {
            this.props.history.push('/serviceBulletins');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ServiceBulletin</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> bulletinNumber: </label>
                                                <input placeholder="bulletinNumber" name="bulletinNumber" className="form-control" value={this.state.bulletinNumber} onChange={this.changebulletinNumberHandler}/>

                                            <label> Category: </label>
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
                                        <button className="btn btn-success" onClick={this.updateServiceBulletin}>Save</button>
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

export default UpdateServiceBulletinComponent
