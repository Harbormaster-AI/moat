import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PublisherService } from '../../../services/Publisher.service';
import { Publisher } from '../../../models/Publisher';
import { SubBaseComponent } from '../../Publisher/sub.base.component';

@Component({
    selector: 'app-create-publisher',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePublisherComponent extends SubBaseComponent implements OnInit {

    title = 'Add Publisher';

    publisherForm: FormGroup;
    publisher: Publisher;

    constructor( http: HttpClient,
        private publisherService: PublisherService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.publisherForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      website: ['', Validators.required],
      InventorySources: ['', ],
      Deals: ['', ],
      CreativeApprovals: ['', ],
      InsertionOrders: ['', ],
      RateCards: ['', ],
      PublisherType: ['', ]
        });
    }

    
    addPublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType): void {
        this.publisherService
        .addPublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType)
            .subscribe(() => {
                this.router.navigate(['/indexPublisher']);
            });
    }

    ngOnInit(): void {
    }
}