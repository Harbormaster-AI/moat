import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AlertService } from '../../../services/Alert.service';
import { SubBaseComponent } from '../../Alert/sub.base.component';


@Component({
    selector: 'app-edit-alert',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAlertComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Alert';

    alertForm: FormGroup;
    alert: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AlertService,
        private fb: FormBuilder
) {
        super(http);
        this.alertForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  title: ['', Validators.required],
      createdAt: ['', Validators.required],
      Metric: ['', ],
      Dashboard: ['', ],
      Dataset: ['', ],
      Rule: ['', ],
      Anomalies: ['', ],
      Subscribers: ['', ],
      Severity: ['', ],
      Status: ['', ]
        });
    }

    
    updateAlert(title, createdAt, Metric, Dashboard, Dataset, Rule, Anomalies, Subscribers, Severity, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAlert(title, createdAt, Metric, Dashboard, Dataset, Rule, Anomalies, Subscribers, Severity, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAlert']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAlert(params['id']).subscribe(res => {
                this.alert = res;
            });
        });
    }
}