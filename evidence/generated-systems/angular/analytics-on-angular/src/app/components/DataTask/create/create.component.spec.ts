
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataTaskComponent } from './create.component';
import { DataTaskService } from '../../../services/DataTask.service';
import { Router } from '@angular/router';

describe('CreateDataTaskComponent', () => {
  let component: CreateDataTaskComponent;
  let fixture: ComponentFixture<CreateDataTaskComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataTaskComponent
      ],
      providers: [
        DataTaskService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataTaskComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});