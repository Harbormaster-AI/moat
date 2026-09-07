import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SubscriberService } from '../../../services/Subscriber.service';
import { Subscriber } from '../../../models/Subscriber';
import { SubBaseComponent } from '../../Subscriber/sub.base.component';

@Component({
    selector: 'app-create-subscriber',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSubscriberComponent extends SubBaseComponent implements OnInit {

    title = 'Add Subscriber';

    subscriberForm: FormGroup;
    subscriber: Subscriber;

    constructor( http: HttpClient,
        private subscriberService: SubscriberService,
        private fb: FormBuilder,
        private router: Router
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

    
    addSubscriber(name, address, Alerts, Channel): void {
        this.subscriberService
        .addSubscriber(name, address, Alerts, Channel)
            .subscribe(() => {
                this.router.navigate(['/indexSubscriber']);
            });
    }

    ngOnInit(): void {
    }
}