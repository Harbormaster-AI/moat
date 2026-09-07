import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataSetService } from '../../../services/DataSet.service';
import { SubBaseComponent } from '../../DataSet/sub.base.component';


@Component({
    selector: 'app-edit-dataSet',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataSetComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataSet';

    dataSetForm: FormGroup;
    dataSet: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataSetService,
        private fb: FormBuilder
) {
        super(http);
        this.dataSetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      schemaVersion: ['', Validators.required],
      refreshSchedule: ['', Validators.required],
      Sensitive: ['', Validators.required],
      Workspace: ['', ],
      Sources: ['', ],
      Pipelines: ['', ],
      SemanticModels: ['', ],
      Dimensions: ['', ],
      Measures: ['', ],
      Metrics: ['', ],
      QualityRules: ['', ],
      LineageNode: ['', ],
      Tags: ['', ],
      DataFormat: ['', ]
        });
    }

    
    updateDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataSet(name, schemaVersion, refreshSchedule, Sensitive, Workspace, Sources, Pipelines, SemanticModels, Dimensions, Measures, Metrics, QualityRules, LineageNode, Tags, DataFormat, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataSet']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataSet(params['id']).subscribe(res => {
                this.dataSet = res;
            });
        });
    }
}