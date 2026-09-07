
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOpportunityLineItemComponent } from './index.component';
import { OpportunityLineItemService } from '../../../services/OpportunityLineItem.service';

describe('IndexOpportunityLineItemComponent', () => {
  let component: IndexOpportunityLineItemComponent;
  let fixture: ComponentFixture<IndexOpportunityLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOpportunityLineItemComponent
      ],
      providers: [
        OpportunityLineItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOpportunityLineItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});