
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBIQueryComponent } from './create.component';
import { BIQueryService } from '../../../services/BIQuery.service';
import { Router } from '@angular/router';

describe('CreateBIQueryComponent', () => {
  let component: CreateBIQueryComponent;
  let fixture: ComponentFixture<CreateBIQueryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBIQueryComponent
      ],
      providers: [
        BIQueryService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBIQueryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});