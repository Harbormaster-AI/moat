
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexAgreementComponent } from './index.component';
import { AgreementService } from '../../../services/Agreement.service';

describe('IndexAgreementComponent', () => {
  let component: IndexAgreementComponent;
  let fixture: ComponentFixture<IndexAgreementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexAgreementComponent
      ],
      providers: [
        AgreementService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexAgreementComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});