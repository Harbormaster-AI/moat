import React, { Component } from 'react'
import APUService from '../services/APUService';

class UpdateAPUComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                model: ''
        }
        this.updateAPU = this.updateAPU.bind(this);

        this.changemodelHandler = this.changemodelHandler.bind(this);
    }

    componentDidMount(){
        APUService.getAPUById(this.state.id).then( (res) =>{
            let aPU = res.data;
            this.setState({
                model: aPU.model
            });
        });
    }

    updateAPU = (e) => {
        e.preventDefault();
        let aPU = {
            aPUId: this.state.id,
            model: this.state.model
        };
        console.log('aPU => ' + JSON.stringify(aPU));
        console.log('id => ' + JSON.stringify(this.state.id));
        APUService.updateAPU(aPU).then( res => {
            this.props.history.push('/aPUs');
        });
    }

    changemodelHandler= (event) => {
        this.setState({model: event.target.value});
    }

    cancel(){
        this.props.history.push('/aPUs');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update APU</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> model: </label>
                                                <input placeholder="model" name="model" className="form-control" value={this.state.model} onChange={this.changemodelHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAPU}>Save</button>
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

export default UpdateAPUComponent
