import React, { Component } from 'react'
import DimensionService from '../services/DimensionService';

class CreateDimensionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                typeTime: '',
                dimensionType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changetypeTimeHandler = this.changetypeTimeHandler.bind(this);
        this.changeDimensionTypeHandler = this.changeDimensionTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DimensionService.getDimensionById(this.state.id).then( (res) =>{
                let dimension = res.data;
                this.setState({
                    name: dimension.name,
                    typeTime: dimension.typeTime,
                    dimensionType: dimension.dimensionType
                });
            });
        }        
    }
    saveOrUpdateDimension = (e) => {
        e.preventDefault();
        let dimension = {
                dimensionId: this.state.id,
                name: this.state.name,
                typeTime: this.state.typeTime,
                dimensionType: this.state.dimensionType
            };
        console.log('dimension => ' + JSON.stringify(dimension));

        // step 5
        if(this.state.id === '_add'){
            dimension.dimensionId=''
            DimensionService.createDimension(dimension).then(res =>{
                this.props.history.push('/dimensions');
            });
        }else{
            DimensionService.updateDimension(dimension).then( res => {
                this.props.history.push('/dimensions');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changetypeTimeHandler= (event) => {
        this.setState({typeTime: event.target.value});
    }
    changeDimensionTypeHandler= (event) => {
        this.setState({dimensionType: event.target.value});
    }

    cancel(){
        this.props.history.push('/dimensions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Dimension</h3>
        }else{
            return <h3 className="text-center">Update Dimension</h3>
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

                                            <label> typeTime:&emsp; </label>
                                                <input type="checkbox" placeholder="typeTime" name="typeTime" className="form-control" value={this.state.typeTime} onChange={this.changetypeTimeHandler}/>


                                            <label> DimensionType:&emsp; </label>
                                                <select value={this.state.dimensionType} onChange={this.changeDimensionTypeHandler}>
                      <option name="DimensionType" className="form-control" >
                          Categorical
                      </option>
                      <option name="DimensionType" className="form-control" >
                          Temporal
                      </option>
                      <option name="DimensionType" className="form-control" >
                          Geospatial
                      </option>
                      <option name="DimensionType" className="form-control" >
                          Hierarchical
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDimension}>Save</button>
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

export default CreateDimensionComponent
