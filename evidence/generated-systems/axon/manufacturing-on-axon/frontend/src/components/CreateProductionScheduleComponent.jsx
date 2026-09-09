import React, { Component } from 'react'
import ProductionScheduleService from '../services/ProductionScheduleService';

class CreateProductionScheduleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                scheduleNumber: '',
                horizonStart: '',
                horizonEnd: '',
                status: ''
        }
        this.changescheduleNumberHandler = this.changescheduleNumberHandler.bind(this);
        this.changehorizonStartHandler = this.changehorizonStartHandler.bind(this);
        this.changehorizonEndHandler = this.changehorizonEndHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ProductionScheduleService.getProductionScheduleById(this.state.id).then( (res) =>{
                let productionSchedule = res.data;
                this.setState({
                    scheduleNumber: productionSchedule.scheduleNumber,
                    horizonStart: productionSchedule.horizonStart,
                    horizonEnd: productionSchedule.horizonEnd,
                    status: productionSchedule.status
                });
            });
        }        
    }
    saveOrUpdateProductionSchedule = (e) => {
        e.preventDefault();
        let productionSchedule = {
                productionScheduleId: this.state.id,
                scheduleNumber: this.state.scheduleNumber,
                horizonStart: this.state.horizonStart,
                horizonEnd: this.state.horizonEnd,
                status: this.state.status
            };
        console.log('productionSchedule => ' + JSON.stringify(productionSchedule));

        // step 5
        if(this.state.id === '_add'){
            productionSchedule.productionScheduleId=''
            ProductionScheduleService.createProductionSchedule(productionSchedule).then(res =>{
                this.props.history.push('/productionSchedules');
            });
        }else{
            ProductionScheduleService.updateProductionSchedule(productionSchedule).then( res => {
                this.props.history.push('/productionSchedules');
            });
        }
    }
    
    changescheduleNumberHandler= (event) => {
        this.setState({scheduleNumber: event.target.value});
    }
    changehorizonStartHandler= (event) => {
        this.setState({horizonStart: event.target.value});
    }
    changehorizonEndHandler= (event) => {
        this.setState({horizonEnd: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/productionSchedules');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ProductionSchedule</h3>
        }else{
            return <h3 className="text-center">Update ProductionSchedule</h3>
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
                                            <label> scheduleNumber:&emsp; </label>
                                                <input placeholder="scheduleNumber" name="scheduleNumber" className="form-control" value={this.state.scheduleNumber} onChange={this.changescheduleNumberHandler}/>

                                            <label> horizonStart:&emsp; </label>
                                                <input type="date" placeholder="horizonStart" name="horizonStart" className="form-control" value={this.state.horizonStart} onChange={this.changehorizonStartHandler}/>

                                            <label> horizonEnd:&emsp; </label>
                                                <input type="date" placeholder="horizonEnd" name="horizonEnd" className="form-control" value={this.state.horizonEnd} onChange={this.changehorizonEndHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Approved
                      </option>
                      <option name="Status" className="form-control" >
                          Frozen
                      </option>
                      <option name="Status" className="form-control" >
                          Completed
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateProductionSchedule}>Save</button>
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

export default CreateProductionScheduleComponent
