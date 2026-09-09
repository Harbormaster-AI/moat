import React, { Component } from 'react'
import GeoRegionService from '../services/GeoRegionService';

class UpdateGeoRegionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                name: '',
                regionType: ''
        }
        this.updateGeoRegion = this.updateGeoRegion.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeRegionTypeHandler = this.changeRegionTypeHandler.bind(this);
    }

    componentDidMount(){
        GeoRegionService.getGeoRegionById(this.state.id).then( (res) =>{
            let geoRegion = res.data;
            this.setState({
                code: geoRegion.code,
                name: geoRegion.name,
                regionType: geoRegion.regionType
            });
        });
    }

    updateGeoRegion = (e) => {
        e.preventDefault();
        let geoRegion = {
            geoRegionId: this.state.id,
            code: this.state.code,
            name: this.state.name,
            regionType: this.state.regionType
        };
        console.log('geoRegion => ' + JSON.stringify(geoRegion));
        console.log('id => ' + JSON.stringify(this.state.id));
        GeoRegionService.updateGeoRegion(geoRegion).then( res => {
            this.props.history.push('/geoRegions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update GeoRegion</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> RegionType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateGeoRegion}>Save</button>
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

export default UpdateGeoRegionComponent
