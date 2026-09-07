
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDirectDebitMandateComponent } from './index.component';
import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';

describe('IndexDirectDebitMandateComponent', () => {
  let component: IndexDirectDebitMandateComponent;
  let fixture: ComponentFixture<IndexDirectDebitMandateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDirectDebitMandateComponent
      ],
      providers: [
        DirectDebitMandateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDirectDebitMandateComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});