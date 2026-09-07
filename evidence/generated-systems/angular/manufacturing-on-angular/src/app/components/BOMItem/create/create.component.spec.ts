
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBOMItemComponent } from './create.component';
import { BOMItemService } from '../../../services/BOMItem.service';
import { Router } from '@angular/router';

describe('CreateBOMItemComponent', () => {
  let component: CreateBOMItemComponent;
  let fixture: ComponentFixture<CreateBOMItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBOMItemComponent
      ],
      providers: [
        BOMItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBOMItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});