import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PublisherService } from '../../../services/Publisher.service';
import { SubBaseComponent } from '../../Publisher/sub.base.component';


@Component({
    selector: 'app-edit-publisher',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPublisherComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Publisher';

    publisherForm: FormGroup;
    publisher: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PublisherService,
        private fb: FormBuilder
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

    
    updatePublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePublisher(name, website, InventorySources, Deals, CreativeApprovals, InsertionOrders, RateCards, PublisherType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPublisher']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPublisher(params['id']).subscribe(res => {
                this.publisher = res;
            });
        });
    }
}