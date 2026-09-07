
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexOpportunityComponent } from './index.component';
import { OpportunityService } from '../../../services/Opportunity.service';

describe('IndexOpportunityComponent', () => {
  let component: IndexOpportunityComponent;
  let fixture: ComponentFixture<IndexOpportunityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexOpportunityComponent
      ],
      providers: [
        OpportunityService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexOpportunityComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});