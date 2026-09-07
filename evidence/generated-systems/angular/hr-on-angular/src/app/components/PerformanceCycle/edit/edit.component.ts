import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PerformanceCycleService } from '../../../services/PerformanceCycle.service';
import { SubBaseComponent } from '../../PerformanceCycle/sub.base.component';


@Component({
    selector: 'app-edit-performanceCycle',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPerformanceCycleComponent extends SubBaseComponent implements OnInit {

    title = 'Edit PerformanceCycle';

    performanceCycleForm: FormGroup;
    performanceCycle: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PerformanceCycleService,
        private fb: FormBuilder
) {
        super(http);
        this.performanceCycleForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      startDate: ['', Validators.required],
      endDate: ['', Validators.required],
      Organization: ['', ],
      Reviews: ['', ],
      Goals: ['', ],
      Status: ['', ]
        });
    }

    
    updatePerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePerformanceCycle(name, startDate, endDate, Organization, Reviews, Goals, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPerformanceCycle']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPerformanceCycle(params['id']).subscribe(res => {
                this.performanceCycle = res;
            });
        });
    }
}