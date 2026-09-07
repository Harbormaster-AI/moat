import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DataCategoryService } from '../../../services/DataCategory.service';
import { SubBaseComponent } from '../../DataCategory/sub.base.component';


@Component({
    selector: 'app-edit-dataCategory',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDataCategoryComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DataCategory';

    dataCategoryForm: FormGroup;
    dataCategory: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DataCategoryService,
        private fb: FormBuilder
) {
        super(http);
        this.dataCategoryForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      description: ['', Validators.required],
      ProcessingActivities: ['', ],
      Records: ['', ],
      DataBreaches: ['', ],
      Classification: ['', ]
        });
    }

    
    updateDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDataCategory(name, description, ProcessingActivities, Records, DataBreaches, Classification, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDataCategory']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDataCategory(params['id']).subscribe(res => {
                this.dataCategory = res;
            });
        });
    }
}