import React, { Component } from 'react'
import SalesRegionService from '../services/SalesRegionService';

class CreateSalesRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                regionCode: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeregionCodeHandler = this.changeregionCodeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            SalesRegionService.getSalesRegionById(this.state.id).then( (res) =>{
                let salesRegion = res.data;
                this.setState({
                    name: salesRegion.name,
                    regionCode: salesRegion.regionCode
                });
            });
        }        
    }
    saveOrUpdateSalesRegion = (e) => {
        e.preventDefault();
        let salesRegion = {
                salesRegionId: this.state.id,
                name: this.state.name,
                regionCode: this.state.regionCode
            };
        console.log('salesRegion => ' + JSON.stringify(salesRegion));

        // step 5
        if(this.state.id === '_add'){
            salesRegion.salesRegionId=''
            SalesRegionService.createSalesRegion(salesRegion).then(res =>{
                this.props.history.push('/salesRegions');
            });
        }else{
            SalesRegionService.updateSalesRegion(salesRegion).then( res => {
                this.props.history.push('/salesRegions');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add SalesRegion</h3>
        }else{
            return <h3 className="text-center">Update SalesRegion</h3>
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

                                            <label> regionCode:&emsp; </label>
                                                <input placeholder="regionCode" name="regionCode" className="form-control" value={this.state.regionCode} onChange={this.changeregionCodeHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateSalesRegion}>Save</button>
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

export default CreateSalesRegionComponent
