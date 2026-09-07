
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBOMComponent } from './create.component';
import { BOMService } from '../../../services/BOM.service';
import { Router } from '@angular/router';

describe('CreateBOMComponent', () => {
  let component: CreateBOMComponent;
  let fixture: ComponentFixture<CreateBOMComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBOMComponent
      ],
      providers: [
        BOMService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBOMComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});