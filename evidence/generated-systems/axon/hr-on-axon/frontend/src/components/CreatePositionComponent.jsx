import React, { Component } from 'react'
import PositionService from '../services/PositionService';

class CreatePositionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                positionCode: '',
                fte: '',
                status: '',
                workLocationType: ''
        }
        this.changepositionCodeHandler = this.changepositionCodeHandler.bind(this);
        this.changefteHandler = this.changefteHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
        this.changeWorkLocationTypeHandler = this.changeWorkLocationTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            PositionService.getPositionById(this.state.id).then( (res) =>{
                let position = res.data;
                this.setState({
                    positionCode: position.positionCode,
                    fte: position.fte,
                    status: position.status,
                    workLocationType: position.workLocationType
                });
            });
        }        
    }
    saveOrUpdatePosition = (e) => {
        e.preventDefault();
        let position = {
                positionId: this.state.id,
                positionCode: this.state.positionCode,
                fte: this.state.fte,
                status: this.state.status,
                workLocationType: this.state.workLocationType
            };
        console.log('position => ' + JSON.stringify(position));

        // step 5
        if(this.state.id === '_add'){
            position.positionId=''
            PositionService.createPosition(position).then(res =>{
                this.props.history.push('/positions');
            });
        }else{
            PositionService.updatePosition(position).then( res => {
                this.props.history.push('/positions');
            });
        }
    }
    
    changepositionCodeHandler= (event) => {
        this.setState({positionCode: event.target.value});
    }
    changefteHandler= (event) => {
        this.setState({fte: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }
    changeWorkLocationTypeHandler= (event) => {
        this.setState({workLocationType: event.target.value});
    }

    cancel(){
        this.props.history.push('/positions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Position</h3>
        }else{
            return <h3 className="text-center">Update Position</h3>
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
                                            <label> positionCode:&emsp; </label>
                                                <input placeholder="positionCode" name="positionCode" className="form-control" value={this.state.positionCode} onChange={this.changepositionCodeHandler}/>

                                            <label> fte:&emsp; </label>
                                                <input placeholder="fte" name="fte" className="form-control" value={this.state.fte} onChange={this.changefteHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Open
                      </option>
                      <option name="Status" className="form-control" >
                          Filled
                      </option>
                      <option name="Status" className="form-control" >
                          Frozen
                      </option>
                      <option name="Status" className="form-control" >
                          Closed
                      </option>
                    </select>

                                            <label> WorkLocationType:&emsp; </label>
                                                <select value={this.state.workLocationType} onChange={this.changeWorkLocationTypeHandler}>
                      <option name="WorkLocationType" className="form-control" >
                          Onsite
                      </option>
                      <option name="WorkLocationType" className="form-control" >
                          Hybrid
                      </option>
                      <option name="WorkLocationType" className="form-control" >
                          Remote
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdatePosition}>Save</button>
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

export default CreatePositionComponent
