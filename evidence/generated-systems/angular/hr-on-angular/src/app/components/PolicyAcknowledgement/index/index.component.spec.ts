
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPolicyAcknowledgementComponent } from './index.component';
import { PolicyAcknowledgementService } from '../../../services/PolicyAcknowledgement.service';

describe('IndexPolicyAcknowledgementComponent', () => {
  let component: IndexPolicyAcknowledgementComponent;
  let fixture: ComponentFixture<IndexPolicyAcknowledgementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPolicyAcknowledgementComponent
      ],
      providers: [
        PolicyAcknowledgementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPolicyAcknowledgementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});