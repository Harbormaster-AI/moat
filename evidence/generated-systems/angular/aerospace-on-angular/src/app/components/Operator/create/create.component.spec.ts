
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateOperatorComponent } from './create.component';
import { OperatorService } from '../../../services/Operator.service';
import { Router } from '@angular/router';

describe('CreateOperatorComponent', () => {
  let component: CreateOperatorComponent;
  let fixture: ComponentFixture<CreateOperatorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateOperatorComponent
      ],
      providers: [
        OperatorService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateOperatorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});