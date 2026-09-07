import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SubscriberService } from '../../../services/Subscriber.service';
import { SubBaseComponent } from '../../Subscriber/sub.base.component';


@Component({
    selector: 'app-edit-subscriber',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSubscriberComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Subscriber';

    subscriberForm: FormGroup;
    subscriber: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SubscriberService,
        private fb: FormBuilder
) {
        super(http);
        this.subscriberForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      address: ['', Validators.required],
      Alerts: ['', ],
      Channel: ['', ]
        });
    }

    
    updateSubscriber(name, address, Alerts, Channel): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSubscriber(name, address, Alerts, Channel, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSubscriber']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSubscriber(params['id']).subscribe(res => {
                this.subscriber = res;
            });
        });
    }
}