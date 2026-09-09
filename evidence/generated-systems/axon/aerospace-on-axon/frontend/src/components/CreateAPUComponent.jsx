import React, { Component } from 'react'
import APUService from '../services/APUService';

class CreateAPUComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                model: ''
        }
        this.changemodelHandler = this.changemodelHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            APUService.getAPUById(this.state.id).then( (res) =>{
                let aPU = res.data;
                this.setState({
                    model: aPU.model
                });
            });
        }        
    }
    saveOrUpdateAPU = (e) => {
        e.preventDefault();
        let aPU = {
                aPUId: this.state.id,
                model: this.state.model
            };
        console.log('aPU => ' + JSON.stringify(aPU));

        // step 5
        if(this.state.id === '_add'){
            aPU.aPUId=''
            APUService.createAPU(aPU).then(res =>{
                this.props.history.push('/aPUs');
            });
        }else{
            APUService.updateAPU(aPU).then( res => {
                this.props.history.push('/aPUs');
            });
        }
    }
    
    changemodelHandler= (event) => {
        this.setState({model: event.target.value});
    }

    cancel(){
        this.props.history.push('/aPUs');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add APU</h3>
        }else{
            return <h3 className="text-center">Update APU</h3>
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
                                            <label> model:&emsp; </label>
                                                <input placeholder="model" name="model" className="form-control" value={this.state.model} onChange={this.changemodelHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAPU}>Save</button>
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

export default CreateAPUComponent
