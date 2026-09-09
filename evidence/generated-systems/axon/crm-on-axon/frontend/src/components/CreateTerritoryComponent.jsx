import React, { Component } from 'react'
import TerritoryService from '../services/TerritoryService';

class CreateTerritoryComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                region: '',
                territoryType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
        this.changeTerritoryTypeHandler = this.changeTerritoryTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TerritoryService.getTerritoryById(this.state.id).then( (res) =>{
                let territory = res.data;
                this.setState({
                    name: territory.name,
                    region: territory.region,
                    territoryType: territory.territoryType
                });
            });
        }        
    }
    saveOrUpdateTerritory = (e) => {
        e.preventDefault();
        let territory = {
                territoryId: this.state.id,
                name: this.state.name,
                region: this.state.region,
                territoryType: this.state.territoryType
            };
        console.log('territory => ' + JSON.stringify(territory));

        // step 5
        if(this.state.id === '_add'){
            territory.territoryId=''
            TerritoryService.createTerritory(territory).then(res =>{
                this.props.history.push('/territorys');
            });
        }else{
            TerritoryService.updateTerritory(territory).then( res => {
                this.props.history.push('/territorys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Territory</h3>
        }else{
            return <h3 className="text-center">Update Territory</h3>
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

                                            <label> region:&emsp; </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                            <label> TerritoryType:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTerritory}>Save</button>
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

export default CreateTerritoryComponent
