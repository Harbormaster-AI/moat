import React, { Component } from 'react'
import SalesRegionService from '../services/SalesRegionService';

class UpdateSalesRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                regionCode: ''
        }
        this.updateSalesRegion = this.updateSalesRegion.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeregionCodeHandler = this.changeregionCodeHandler.bind(this);
    }

    componentDidMount(){
        SalesRegionService.getSalesRegionById(this.state.id).then( (res) =>{
            let salesRegion = res.data;
            this.setState({
                name: salesRegion.name,
                regionCode: salesRegion.regionCode
            });
        });
    }

    updateSalesRegion = (e) => {
        e.preventDefault();
        let salesRegion = {
            salesRegionId: this.state.id,
            name: this.state.name,
            regionCode: this.state.regionCode
        };
        console.log('salesRegion => ' + JSON.stringify(salesRegion));
        console.log('id => ' + JSON.stringify(this.state.id));
        SalesRegionService.updateSalesRegion(salesRegion).then( res => {
            this.props.history.push('/salesRegions');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeregionCodeHandler= (event) => {
        this.setState({regionCode: event.target.value});
    }

    cancel(){
        this.props.history.push('/salesRegions');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update SalesRegion</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> regionCode: </label>
                                                <input placeholder="regionCode" name="regionCode" className="form-control" value={this.state.regionCode} onChange={this.changeregionCodeHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateSalesRegion}>Save</button>
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

export default UpdateSalesRegionComponent
