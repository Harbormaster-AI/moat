import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TagService } from '../../../services/Tag.service';
import { SubBaseComponent } from '../../Tag/sub.base.component';


@Component({
    selector: 'app-edit-tag',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTagComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Tag';

    tagForm: FormGroup;
    tag: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TagService,
        private fb: FormBuilder
) {
        super(http);
        this.tagForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      Datasets: ['', ],
      Models: ['', ],
      ModelVersions: ['', ],
      Dashboards: ['', ],
      Reports: ['', ],
      FeatureSets: ['', ],
      Metrics: ['', ],
      Category: ['', ]
        });
    }

    
    updateTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTag(name, Datasets, Models, ModelVersions, Dashboards, Reports, FeatureSets, Metrics, Category, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTag']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTag(params['id']).subscribe(res => {
                this.tag = res;
            });
        });
    }
}