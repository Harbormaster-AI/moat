
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataProviderComponent } from './create.component';
import { DataProviderService } from '../../../services/DataProvider.service';
import { Router } from '@angular/router';

describe('CreateDataProviderComponent', () => {
  let component: CreateDataProviderComponent;
  let fixture: ComponentFixture<CreateDataProviderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataProviderComponent
      ],
      providers: [
        DataProviderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataProviderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});