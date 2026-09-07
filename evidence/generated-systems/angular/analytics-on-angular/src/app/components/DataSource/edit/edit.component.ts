import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataSourceService } from '../../../services/DataSource.service';
import { SubBaseComponent } from '../../DataSource/sub.base.component';


@Component({
    selector: 'app-edit-dataSource',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataSourceComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataSource';

    dataSourceForm: FormGroup;
    dataSource: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataSourceService,
        private fb: FormBuilder
) {
        super(http);
        this.dataSourceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      connection: ['', Validators.required],
      Streaming: ['', Validators.required],
      Workspace: ['', ],
      ProducedDatasets: ['', ],
      Pipelines: ['', ],
      SourceType: ['', ],
      Format: ['', ]
        });
    }

    
    updateDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataSource(name, connection, Streaming, Workspace, ProducedDatasets, Pipelines, SourceType, Format, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataSource']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataSource(params['id']).subscribe(res => {
                this.dataSource = res;
            });
        });
    }
}