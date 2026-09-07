
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataSetComponent } from './create.component';
import { DataSetService } from '../../../services/DataSet.service';
import { Router } from '@angular/router';

describe('CreateDataSetComponent', () => {
  let component: CreateDataSetComponent;
  let fixture: ComponentFixture<CreateDataSetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataSetComponent
      ],
      providers: [
        DataSetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataSetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});