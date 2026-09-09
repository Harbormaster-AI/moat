import React, { Component } from 'react'
import GeoRegionService from '../services/GeoRegionService';

class CreateGeoRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                name: '',
                regionType: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeRegionTypeHandler = this.changeRegionTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            GeoRegionService.getGeoRegionById(this.state.id).then( (res) =>{
                let geoRegion = res.data;
                this.setState({
                    code: geoRegion.code,
                    name: geoRegion.name,
                    regionType: geoRegion.regionType
                });
            });
        }        
    }
    saveOrUpdateGeoRegion = (e) => {
        e.preventDefault();
        let geoRegion = {
                geoRegionId: this.state.id,
                code: this.state.code,
                name: this.state.name,
                regionType: this.state.regionType
            };
        console.log('geoRegion => ' + JSON.stringify(geoRegion));

        // step 5
        if(this.state.id === '_add'){
            geoRegion.geoRegionId=''
            GeoRegionService.createGeoRegion(geoRegion).then(res =>{
                this.props.history.push('/geoRegions');
            });
        }else{
            GeoRegionService.updateGeoRegion(geoRegion).then( res => {
                this.props.history.push('/geoRegions');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeRegionTypeHandler= (event) => {
        this.setState({regionType: event.target.value});
    }

    cancel(){
        this.props.history.push('/geoRegions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add GeoRegion</h3>
        }else{
            return <h3 className="text-center">Update GeoRegion</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> RegionType:&emsp; </label>
                                                <select value={this.state.regionType} onChange={this.changeRegionTypeHandler}>
                      <option name="RegionType" className="form-control" >
                          Country
                      </option>
                      <option name="RegionType" className="form-control" >
                          State
                      </option>
                      <option name="RegionType" className="form-control" >
                          Province
                      </option>
                      <option name="RegionType" className="form-control" >
                          City
                      </option>
                      <option name="RegionType" className="form-control" >
                          DMA
                      </option>
                      <option name="RegionType" className="form-control" >
                          PostalCode
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateGeoRegion}>Save</button>
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

export default CreateGeoRegionComponent
