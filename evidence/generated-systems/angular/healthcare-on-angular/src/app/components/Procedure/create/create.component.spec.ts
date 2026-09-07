
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateProcedureComponent } from './create.component';
import { ProcedureService } from '../../../services/Procedure.service';
import { Router } from '@angular/router';

describe('CreateProcedureComponent', () => {
  let component: CreateProcedureComponent;
  let fixture: ComponentFixture<CreateProcedureComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateProcedureComponent
      ],
      providers: [
        ProcedureService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateProcedureComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});