import React, { Component } from 'react'
import AvionicsSuiteService from '../services/AvionicsSuiteService';

class CreateAvionicsSuiteComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                suiteName: '',
                softwareBaseline: ''
        }
        this.changesuiteNameHandler = this.changesuiteNameHandler.bind(this);
        this.changesoftwareBaselineHandler = this.changesoftwareBaselineHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AvionicsSuiteService.getAvionicsSuiteById(this.state.id).then( (res) =>{
                let avionicsSuite = res.data;
                this.setState({
                    suiteName: avionicsSuite.suiteName,
                    softwareBaseline: avionicsSuite.softwareBaseline
                });
            });
        }        
    }
    saveOrUpdateAvionicsSuite = (e) => {
        e.preventDefault();
        let avionicsSuite = {
                avionicsSuiteId: this.state.id,
                suiteName: this.state.suiteName,
                softwareBaseline: this.state.softwareBaseline
            };
        console.log('avionicsSuite => ' + JSON.stringify(avionicsSuite));

        // step 5
        if(this.state.id === '_add'){
            avionicsSuite.avionicsSuiteId=''
            AvionicsSuiteService.createAvionicsSuite(avionicsSuite).then(res =>{
                this.props.history.push('/avionicsSuites');
            });
        }else{
            AvionicsSuiteService.updateAvionicsSuite(avionicsSuite).then( res => {
                this.props.history.push('/avionicsSuites');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AvionicsSuite</h3>
        }else{
            return <h3 className="text-center">Update AvionicsSuite</h3>
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
                                            <label> suiteName:&emsp; </label>
                                                <input placeholder="suiteName" name="suiteName" className="form-control" value={this.state.suiteName} onChange={this.changesuiteNameHandler}/>

                                            <label> softwareBaseline:&emsp; </label>
                                                <input placeholder="softwareBaseline" name="softwareBaseline" className="form-control" value={this.state.softwareBaseline} onChange={this.changesoftwareBaselineHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAvionicsSuite}>Save</button>
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

export default CreateAvionicsSuiteComponent
