
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataSourceComponent } from './create.component';
import { DataSourceService } from '../../../services/DataSource.service';
import { Router } from '@angular/router';

describe('CreateDataSourceComponent', () => {
  let component: CreateDataSourceComponent;
  let fixture: ComponentFixture<CreateDataSourceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataSourceComponent
      ],
      providers: [
        DataSourceService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataSourceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});