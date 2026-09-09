import React, { Component } from 'react'
import BOMService from '../services/BOMService';

class CreateBOMComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                bomNumber: '',
                revision: '',
                effectivityStart: '',
                effectivityEnd: '',
                status: ''
        }
        this.changebomNumberHandler = this.changebomNumberHandler.bind(this);
        this.changerevisionHandler = this.changerevisionHandler.bind(this);
        this.changeeffectivityStartHandler = this.changeeffectivityStartHandler.bind(this);
        this.changeeffectivityEndHandler = this.changeeffectivityEndHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            BOMService.getBOMById(this.state.id).then( (res) =>{
                let bOM = res.data;
                this.setState({
                    bomNumber: bOM.bomNumber,
                    revision: bOM.revision,
                    effectivityStart: bOM.effectivityStart,
                    effectivityEnd: bOM.effectivityEnd,
                    status: bOM.status
                });
            });
        }        
    }
    saveOrUpdateBOM = (e) => {
        e.preventDefault();
        let bOM = {
                bOMId: this.state.id,
                bomNumber: this.state.bomNumber,
                revision: this.state.revision,
                effectivityStart: this.state.effectivityStart,
                effectivityEnd: this.state.effectivityEnd,
                status: this.state.status
            };
        console.log('bOM => ' + JSON.stringify(bOM));

        // step 5
        if(this.state.id === '_add'){
            bOM.bOMId=''
            BOMService.createBOM(bOM).then(res =>{
                this.props.history.push('/bOMs');
            });
        }else{
            BOMService.updateBOM(bOM).then( res => {
                this.props.history.push('/bOMs');
            });
        }
    }
    
    changebomNumberHandler= (event) => {
        this.setState({bomNumber: event.target.value});
    }
    changerevisionHandler= (event) => {
        this.setState({revision: event.target.value});
    }
    changeeffectivityStartHandler= (event) => {
        this.setState({effectivityStart: event.target.value});
    }
    changeeffectivityEndHandler= (event) => {
        this.setState({effectivityEnd: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/bOMs');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add BOM</h3>
        }else{
            return <h3 className="text-center">Update BOM</h3>
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
                                            <label> bomNumber:&emsp; </label>
                                                <input placeholder="bomNumber" name="bomNumber" className="form-control" value={this.state.bomNumber} onChange={this.changebomNumberHandler}/>

                                            <label> revision:&emsp; </label>
                                                <input placeholder="revision" name="revision" className="form-control" value={this.state.revision} onChange={this.changerevisionHandler}/>

                                            <label> effectivityStart:&emsp; </label>
                                                <input type="date" placeholder="effectivityStart" name="effectivityStart" className="form-control" value={this.state.effectivityStart} onChange={this.changeeffectivityStartHandler}/>

                                            <label> effectivityEnd:&emsp; </label>
                                                <input type="date" placeholder="effectivityEnd" name="effectivityEnd" className="form-control" value={this.state.effectivityEnd} onChange={this.changeeffectivityEndHandler}/>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Obsolete
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateBOM}>Save</button>
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

export default CreateBOMComponent
