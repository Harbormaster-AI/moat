
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditDirectDebitMandateComponent } from './edit.component';
import { DirectDebitMandateService } from '../../../services/DirectDebitMandate.service';

describe('EditDirectDebitMandateComponent', () => {
  let component: EditDirectDebitMandateComponent;
  let fixture: ComponentFixture<EditDirectDebitMandateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditDirectDebitMandateComponent
      ],
      providers: [
        DirectDebitMandateService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditDirectDebitMandateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});