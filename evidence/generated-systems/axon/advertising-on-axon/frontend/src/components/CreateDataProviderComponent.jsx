import React, { Component } from 'react'
import DataProviderService from '../services/DataProviderService';

class CreateDataProviderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                website: '',
                providerType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            DataProviderService.getDataProviderById(this.state.id).then( (res) =>{
                let dataProvider = res.data;
                this.setState({
                    name: dataProvider.name,
                    website: dataProvider.website,
                    providerType: dataProvider.providerType
                });
            });
        }        
    }
    saveOrUpdateDataProvider = (e) => {
        e.preventDefault();
        let dataProvider = {
                dataProviderId: this.state.id,
                name: this.state.name,
                website: this.state.website,
                providerType: this.state.providerType
            };
        console.log('dataProvider => ' + JSON.stringify(dataProvider));

        // step 5
        if(this.state.id === '_add'){
            dataProvider.dataProviderId=''
            DataProviderService.createDataProvider(dataProvider).then(res =>{
                this.props.history.push('/dataProviders');
            });
        }else{
            DataProviderService.updateDataProvider(dataProvider).then( res => {
                this.props.history.push('/dataProviders');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeProviderTypeHandler= (event) => {
        this.setState({providerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataProviders');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DataProvider</h3>
        }else{
            return <h3 className="text-center">Update DataProvider</h3>
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

                                            <label> website:&emsp; </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> ProviderType:&emsp; </label>
                                                <select value={this.state.providerType} onChange={this.changeProviderTypeHandler}>
                      <option name="ProviderType" className="form-control" >
                          FirstParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          SecondParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          ThirdParty
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDataProvider}>Save</button>
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

export default CreateDataProviderComponent
