
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCreativeAssetComponent } from './create.component';
import { CreativeAssetService } from '../../../services/CreativeAsset.service';
import { Router } from '@angular/router';

describe('CreateCreativeAssetComponent', () => {
  let component: CreateCreativeAssetComponent;
  let fixture: ComponentFixture<CreateCreativeAssetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCreativeAssetComponent
      ],
      providers: [
        CreativeAssetService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCreativeAssetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});