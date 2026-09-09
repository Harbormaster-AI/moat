import React, { Component } from 'react'
import TerritoryService from '../services/TerritoryService';

class UpdateTerritoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                region: '',
                territoryType: ''
        }
        this.updateTerritory = this.updateTerritory.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
        this.changeTerritoryTypeHandler = this.changeTerritoryTypeHandler.bind(this);
    }

    componentDidMount(){
        TerritoryService.getTerritoryById(this.state.id).then( (res) =>{
            let territory = res.data;
            this.setState({
                name: territory.name,
                region: territory.region,
                territoryType: territory.territoryType
            });
        });
    }

    updateTerritory = (e) => {
        e.preventDefault();
        let territory = {
            territoryId: this.state.id,
            name: this.state.name,
            region: this.state.region,
            territoryType: this.state.territoryType
        };
        console.log('territory => ' + JSON.stringify(territory));
        console.log('id => ' + JSON.stringify(this.state.id));
        TerritoryService.updateTerritory(territory).then( res => {
            this.props.history.push('/territorys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeregionHandler= (event) => {
        this.setState({region: event.target.value});
    }
    changeTerritoryTypeHandler= (event) => {
        this.setState({territoryType: event.target.value});
    }

    cancel(){
        this.props.history.push('/territorys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Territory</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> region: </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                            <label> TerritoryType: </label>
                                                <select value={this.state.territoryType} onChange={this.changeTerritoryTypeHandler}>
                      <option name="TerritoryType" className="form-control" >
                          Geographic
                      </option>
                      <option name="TerritoryType" className="form-control" >
                          Industry
                      </option>
                      <option name="TerritoryType" className="form-control" >
                          NamedAccount
                      </option>
                      <option name="TerritoryType" className="form-control" >
                          Segment
                      </option>
                      <option name="TerritoryType" className="form-control" >
                          Hybrid
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTerritory}>Save</button>
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

export default UpdateTerritoryComponent
