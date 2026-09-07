import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BIQueryService } from '../../../services/BIQuery.service';
import { SubBaseComponent } from '../../BIQuery/sub.base.component';


@Component({
    selector: 'app-edit-bIQuery',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBIQueryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BIQuery';

    bIQueryForm: FormGroup;
    bIQuery: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BIQueryService,
        private fb: FormBuilder
) {
        super(http);
        this.bIQueryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      text: ['', Validators.required],
      Workspace: ['', ],
      Datasets: ['', ],
      Reports: ['', ],
      Dashboards: ['', ],
      Notebooks: ['', ],
      Dialect: ['', ]
        });
    }

    
    updateBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBIQuery(name, text, Workspace, Datasets, Reports, Dashboards, Notebooks, Dialect, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBIQuery']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBIQuery(params['id']).subscribe(res => {
                this.bIQuery = res;
            });
        });
    }
}