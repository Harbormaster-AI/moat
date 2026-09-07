
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCreditorComponent } from './create.component';
import { CreditorService } from '../../../services/Creditor.service';
import { Router } from '@angular/router';

describe('CreateCreditorComponent', () => {
  let component: CreateCreditorComponent;
  let fixture: ComponentFixture<CreateCreditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCreditorComponent
      ],
      providers: [
        CreditorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCreditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});