import React, { Component } from 'react'
import DistributorService from '../services/DistributorService';

class UpdateDistributorComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                licenseNumber: '',
                region: '',
                distributorType: ''
        }
        this.updateDistributor = this.updateDistributor.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelicenseNumberHandler = this.changelicenseNumberHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
        this.changeDistributorTypeHandler = this.changeDistributorTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateDistributor = (e) => {
        e.preventDefault();
        let distributor = {
            distributorId: this.state.id,
            name: this.state.name,
            licenseNumber: this.state.licenseNumber,
            region: this.state.region,
            distributorType: this.state.distributorType
        };
        console.log('distributor => ' + JSON.stringify(distributor));
        console.log('id => ' + JSON.stringify(this.state.id));
        DistributorService.updateDistributor(distributor).then( res => {
            this.props.history.push('/distributors');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Distributor</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> licenseNumber: </label>
                                                <input placeholder="licenseNumber" name="licenseNumber" className="form-control" value={this.state.licenseNumber} onChange={this.changelicenseNumberHandler}/>

                                            <label> region: </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                            <label> DistributorType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateDistributor}>Save</button>
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

export default UpdateDistributorComponent
