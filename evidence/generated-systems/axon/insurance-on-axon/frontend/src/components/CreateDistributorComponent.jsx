import React, { Component } from 'react'
import DistributorService from '../services/DistributorService';

class CreateDistributorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                licenseNumber: '',
                region: '',
                distributorType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelicenseNumberHandler = this.changelicenseNumberHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
        this.changeDistributorTypeHandler = this.changeDistributorTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DistributorService.getDistributorById(this.state.id).then( (res) =>{
                let distributor = res.data;
                this.setState({
                    name: distributor.name,
                    licenseNumber: distributor.licenseNumber,
                    region: distributor.region,
                    distributorType: distributor.distributorType
                });
            });
        }        
    }
    saveOrUpdateDistributor = (e) => {
        e.preventDefault();
        let distributor = {
                distributorId: this.state.id,
                name: this.state.name,
                licenseNumber: this.state.licenseNumber,
                region: this.state.region,
                distributorType: this.state.distributorType
            };
        console.log('distributor => ' + JSON.stringify(distributor));

        // step 5
        if(this.state.id === '_add'){
            distributor.distributorId=''
            DistributorService.createDistributor(distributor).then(res =>{
                this.props.history.push('/distributors');
            });
        }else{
            DistributorService.updateDistributor(distributor).then( res => {
                this.props.history.push('/distributors');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelicenseNumberHandler= (event) => {
        this.setState({licenseNumber: event.target.value});
    }
    changeregionHandler= (event) => {
        this.setState({region: event.target.value});
    }
    changeDistributorTypeHandler= (event) => {
        this.setState({distributorType: event.target.value});
    }

    cancel(){
        this.props.history.push('/distributors');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Distributor</h3>
        }else{
            return <h3 className="text-center">Update Distributor</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> licenseNumber:&emsp; </label>
                                                <input placeholder="licenseNumber" name="licenseNumber" className="form-control" value={this.state.licenseNumber} onChange={this.changelicenseNumberHandler}/>

                                            <label> region:&emsp; </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                            <label> DistributorType:&emsp; </label>
                                                <select value={this.state.distributorType} onChange={this.changeDistributorTypeHandler}>
                      <option name="DistributorType" className="form-control" >
                          Agency
                      </option>
                      <option name="DistributorType" className="form-control" >
                          Broker
                      </option>
                      <option name="DistributorType" className="form-control" >
                          Direct
                      </option>
                      <option name="DistributorType" className="form-control" >
                          Bancassurance
                      </option>
                      <option name="DistributorType" className="form-control" >
                          AffinityPartner
                      </option>
                      <option name="DistributorType" className="form-control" >
                          OnlineAggregator
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDistributor}>Save</button>
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

export default CreateDistributorComponent
