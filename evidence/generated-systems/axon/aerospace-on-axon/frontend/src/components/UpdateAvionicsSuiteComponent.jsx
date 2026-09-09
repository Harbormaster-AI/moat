import React, { Component } from 'react'
import AvionicsSuiteService from '../services/AvionicsSuiteService';

class UpdateAvionicsSuiteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                suiteName: '',
                softwareBaseline: ''
        }
        this.updateAvionicsSuite = this.updateAvionicsSuite.bind(this);

        this.changesuiteNameHandler = this.changesuiteNameHandler.bind(this);
        this.changesoftwareBaselineHandler = this.changesoftwareBaselineHandler.bind(this);
    }

    componentDidMount(){
        AvionicsSuiteService.getAvionicsSuiteById(this.state.id).then( (res) =>{
            let avionicsSuite = res.data;
            this.setState({
                suiteName: avionicsSuite.suiteName,
                softwareBaseline: avionicsSuite.softwareBaseline
            });
        });
    }

    updateAvionicsSuite = (e) => {
        e.preventDefault();
        let avionicsSuite = {
            avionicsSuiteId: this.state.id,
            suiteName: this.state.suiteName,
            softwareBaseline: this.state.softwareBaseline
        };
        console.log('avionicsSuite => ' + JSON.stringify(avionicsSuite));
        console.log('id => ' + JSON.stringify(this.state.id));
        AvionicsSuiteService.updateAvionicsSuite(avionicsSuite).then( res => {
            this.props.history.push('/avionicsSuites');
        });
    }

    changesuiteNameHandler= (event) => {
        this.setState({suiteName: event.target.value});
    }
    changesoftwareBaselineHandler= (event) => {
        this.setState({softwareBaseline: event.target.value});
    }

    cancel(){
        this.props.history.push('/avionicsSuites');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AvionicsSuite</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> suiteName: </label>
                                                <input placeholder="suiteName" name="suiteName" className="form-control" value={this.state.suiteName} onChange={this.changesuiteNameHandler}/>

                                            <label> softwareBaseline: </label>
                                                <input placeholder="softwareBaseline" name="softwareBaseline" className="form-control" value={this.state.softwareBaseline} onChange={this.changesoftwareBaselineHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAvionicsSuite}>Save</button>
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

export default UpdateAvionicsSuiteComponent
